module access_sui::token {
    use sui::object::{Self, ID, UID};
    use sui::tx_context::{Self, TxContext};
    use sui::transfer;
    use sui::coin::{Self, Coin};
    use sui::sui::SUI;
    use sui::clock::{Self, Clock};
    use access_sui::events;
    use access_sui::content::{Self, Content};

    /// Represents an AccessToken that grants access to a specific content.
    public struct AccessToken has key, store {
        id: UID,
        owner: address,
        content_id: ID,
        issued_at: u64,
        expires_at: u64,
        revoked: bool,
    }

    /// Buys an AccessToken for a specific content.
    public entry fun buy_access_token(
        content: &Content,
        mut payment: Coin<SUI>,
        duration_seconds: u64,
        clock: &Clock,
        ctx: &mut TxContext
    ) {
        let content_price = content::price(content);
        assert!(coin::value(&payment) >= content_price, 0);

        let sender = tx_context::sender(ctx);
        let current_timestamp = clock::timestamp_ms(clock);
        let expires_at = current_timestamp + duration_seconds * 1000;

        let token = AccessToken {
            id: object::new(ctx),
            owner: sender,
            content_id: object::id(content),
            issued_at: current_timestamp,
            expires_at: expires_at,
            revoked: false,
        };

        transfer::public_transfer(coin::split(&mut payment, content_price, ctx), content::creator(content));
        transfer::public_transfer(payment, sender);

        events::emit_token_issued(object::id(&token), object::id(content), sender, current_timestamp, expires_at);

        transfer::transfer(token, sender);
    }

    /// Transfers an AccessToken to a new owner.
    public entry fun transfer_access_token(
        mut token: AccessToken,
        recipient: address,
        mut payment: Coin<SUI>,
        content: &Content,
        ctx: &mut TxContext
    ) {
        let sender = tx_context::sender(ctx);
        assert!(sender == token.owner, 0);

        let creator_address = content::creator(content);
        let royalty_rate_numerator = 10;
        let royalty_rate_denominator = 100;
        let total_payment = coin::value(&payment);
        let _royalty_amount = total_payment * royalty_rate_numerator / royalty_rate_denominator;

        transfer::public_transfer(coin::split(&mut payment, _royalty_amount, ctx), creator_address);
        transfer::public_transfer(payment, sender);

        token.owner = recipient;
        transfer::transfer(token, recipient);
    }

    /// Revokes an AccessToken.
    public entry fun revoke_access_token(
        mut token: AccessToken,
        content: &Content,
        ctx: &mut TxContext
    ) {
        let sender = tx_context::sender(ctx);
        assert!(sender == content::creator(content), 0);

        token.revoked = true;

        events::emit_token_revoked(object::id(&token), token.content_id, sender);

        transfer::transfer(token, sender);
    }

    /// Checks if an AccessToken is valid for a given content.
    public fun is_access_token_valid(token: &AccessToken, content: &Content, clock: &Clock): bool {
        token.content_id == object::id(content) &&
        !token.revoked &&
        token.expires_at > clock::timestamp_ms(clock)
    }
}