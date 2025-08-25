module access_sui::token {
    use sui::object::{Self, ID, UID};
    use sui::tx_context::{Self, TxContext};
    use sui::transfer;
    use sui::coin::{Self, Coin};
    use sui::sui::SUI;
    use sui::clock::{Self, Clock};
    use access_sui::events;
    use access_sui::content::{Self, Content};

    /// 代表一个AccessToken，它授予对特定内容的访问权限。
    /// 这是一个具有唯一ID的对象，由用户拥有。
    public struct AccessToken has key, store {
        id: UID,
        /// AccessToken的所有者地址。
        owner: address,
        /// 此Token授予访问权限的内容对象的ID。
        content_id: ID,
        /// AccessToken的颁发时间戳。
        issued_at: u64,
        /// AccessToken的过期时间戳（Unix时间戳）。
        expires_at: u64,
        /// 如果Token已被创作者撤销，则为true。
        revoked: bool,
    }

    /// 为特定内容购买AccessToken。
    public entry fun buy_access_token(
        content: &Content,
        mut payment: Coin<SUI>,
        duration_seconds: u64,
        clock: &Clock,
        ctx: &mut TxContext
    ) {
        // 检查付款金额是否足够。
        let content_price = content::price(content);
        assert!(coin::value(&payment) >= content_price, 0);

        let sender = tx_context::sender(ctx);
        let current_timestamp = clock::timestamp_ms(clock);
        let expires_at = current_timestamp + duration_seconds * 1000; // 将秒转换为毫秒

        // 创建一个新的AccessToken对象。
        let token = AccessToken {
            id: object::new(ctx),
            owner: sender,
            content_id: object::id(content),
            issued_at: current_timestamp,
            expires_at: expires_at,
            revoked: false,
        };

        // 将付款转移给内容创作者。
        transfer::public_transfer(coin::split(&mut payment, content_price, ctx), content::creator(content));
        // 将剩余的SUI返还给发送者。
        transfer::public_transfer(payment, sender);

        // 发出AccessToken颁发事件。
        events::emit_token_issued(object::id(&token), object::id(content), sender, current_timestamp, expires_at);

        // 将AccessToken对象转移给购买者。
        transfer::transfer(token, sender);
    }

    /// 将AccessToken转移给新的所有者。
    public entry fun transfer_access_token(
        mut token: AccessToken,
        recipient: address,
        mut payment: Coin<SUI>,
        content: &Content,
        ctx: &mut TxContext
    ) {
        // 只有所有者才能转移Token。
        let sender = tx_context::sender(ctx);
        assert!(sender == token.owner, 0);

        // 计算并转移版税给创作者。
        let creator_address = content::creator(content);
        let royalty_rate_numerator = 10; // 10%的版税率
        let royalty_rate_denominator = 100;
        let total_payment = coin::value(&payment);
        let _royalty_amount = total_payment * royalty_rate_numerator / royalty_rate_denominator;

        transfer::public_transfer(coin::split(&mut payment, _royalty_amount, ctx), creator_address);
        // 将剩余付款转移给原所有者（卖方）。
        transfer::public_transfer(payment, sender);

        // 更改所有者并转移AccessToken对象。
        token.owner = recipient;
        transfer::transfer(token, recipient);
    }

    /// 撤销一个AccessToken。
    public entry fun revoke_access_token(
        mut token: AccessToken,
        content: &Content,
        ctx: &mut TxContext
    ) {
        // 只有内容创作者才能撤销其内容的Token。
        let sender = tx_context::sender(ctx);
        assert!(sender == content::creator(content), 0);

        token.revoked = true;

        // 发出AccessToken撤销事件。
        events::emit_token_revoked(object::id(&token), token.content_id, sender);

        // 将被撤销的Token返回给创作者或销毁它。
        transfer::transfer(token, sender);
    }

    /// 检查AccessToken对于给定内容是否有效。
    public fun is_access_token_valid(token: &AccessToken, content: &Content, clock: &Clock): bool {
        token.content_id == object::id(content) &&
        !token.revoked &&
        token.expires_at > clock::timestamp_ms(clock)
    }
}