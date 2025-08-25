module access_sui::content {
    use sui::object::{Self, UID};
    use sui::tx_context::{Self, TxContext};
    use sui::transfer;
    use sui::table::{Self, Table};
    use std::option::{Self, Option};
    use access_sui::events;

    /// Represents a piece of content published by a creator.
    public struct Content has key, store {
        id: UID,
        creator: address,
        uri: vector<u8>,
        content_hash: vector<u8>,
        price: u64,
        access_levels: Table<vector<u8>, u8>,
    }

    /// Publishes new content.
    public entry fun publish_content(
        creator: address,
        _uri: vector<u8>,
        content_hash: vector<u8>,
        price: u64,
        ctx: &mut TxContext
    ) {
        let mut content = Content {
            id: object::new(ctx),
            creator: creator,
            uri: _uri,
            content_hash: content_hash,
            price: price,
            access_levels: table::new(ctx),
        };

        table::add(&mut content.access_levels, b"basic", 1);

        events::emit_content_published(object::id(&content), creator, _uri, price);

        transfer::transfer(content, creator);
    }

    /// Adds an access level to a content object.
    public entry fun add_access_level(
        content: &mut Content,
        level_name: vector<u8>,
        level_value: u8,
        ctx: &mut TxContext
    ) {
        assert!(tx_context::sender(ctx) == content.creator, 0);
        table::add(&mut content.access_levels, level_name, level_value);
    }

    /// Gets an access level from a content object.
    public fun get_access_level(content: &Content, level_name: vector<u8>): Option<u8> {
        if (table::contains(&content.access_levels, level_name)) {
            option::some<u8>(*table::borrow(&content.access_levels, level_name))
        } else {
            option::none<u8>()
        }
    }

    /// Gets the content hash for a given Content object.
    public fun get_content_hash(content: &Content): vector<u8> {
        content.content_hash
    }

    public fun price(content: &Content): u64 {
        content.price
    }

    public fun creator(content: &Content): address {
        content.creator
    }
}