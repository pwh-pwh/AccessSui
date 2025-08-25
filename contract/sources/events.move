module access_sui::events {
    use sui::object::ID;
    use sui::event;

    /// 发布新内容时发出的事件。
    public struct ContentPublished has copy, drop {
        /// 内容对象的ID。
        content_id: ID,
        /// 内容创作者的地址。
        creator: address,
        /// 内容的URI。
        uri: vector<u8>,
        /// 内容的价格。
        price: u64,
    }

    /// 发出内容发布事件。
    public fun emit_content_published(content_id: ID, creator: address, uri: vector<u8>, price: u64) {
        event::emit(ContentPublished {
            content_id,
            creator,
            uri,
            price,
        })
    }

    /// 颁发AccessToken时发出的事件。
    public struct TokenIssued has copy, drop {
        /// AccessToken对象的ID。
        token_id: ID,
        /// 内容对象的ID。
        content_id: ID,
        /// AccessToken所有者的地址。
        owner: address,
        /// AccessToken的颁发时间戳。
        issued_at: u64,
        /// AccessToken的过期时间戳。
        expires_at: u64,
    }

    /// 发出AccessToken颁发事件。
    public fun emit_token_issued(token_id: ID, content_id: ID, owner: address, issued_at: u64, expires_at: u64) {
        event::emit(TokenIssued {
            token_id,
            content_id,
            owner,
            issued_at,
            expires_at,
        })
    }

    /// 撤销AccessToken时发出的事件。
    public struct TokenRevoked has copy, drop {
        /// AccessToken对象的ID。
        token_id: ID,
        /// 内容对象的ID。
        content_id: ID,
        /// AccessToken所有者的地址。
        owner: address,
    }

    /// 发出AccessToken撤销事件。
    public fun emit_token_revoked(token_id: ID, content_id: ID, owner: address) {
        event::emit(TokenRevoked {
            token_id,
            content_id,
            owner,
        })
    }
}