module access_sui::events {
    use sui::object::ID;
    use sui::event;

    /// Event emitted when new content is published.
    public struct ContentPublished has copy, drop {
        content_id: ID,
        creator: address,
        uri: vector<u8>,
        price: u64,
    }

    public fun emit_content_published(content_id: ID, creator: address, uri: vector<u8>, price: u64) {
        event::emit(ContentPublished {
            content_id,
            creator,
            uri,
            price,
        })
    }

    /// Event emitted when an AccessToken is issued.
    public struct TokenIssued has copy, drop {
        token_id: ID,
        content_id: ID,
        owner: address,
        issued_at: u64,
        expires_at: u64,
    }

    public fun emit_token_issued(token_id: ID, content_id: ID, owner: address, issued_at: u64, expires_at: u64) {
        event::emit(TokenIssued {
            token_id,
            content_id,
            owner,
            issued_at,
            expires_at,
        })
    }

    /// Event emitted when an AccessToken is revoked.
    public struct TokenRevoked has copy, drop {
        token_id: ID,
        content_id: ID,
        owner: address,
    }

    public fun emit_token_revoked(token_id: ID, content_id: ID, owner: address) {
        event::emit(TokenRevoked {
            token_id,
            content_id,
            owner,
        })
    }
}