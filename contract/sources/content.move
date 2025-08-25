module access_sui::content {
    use sui::object::{Self, UID};
    use sui::tx_context::{Self, TxContext};
    use sui::transfer;
    use sui::table::{Self, Table};
    use std::option::{Self, Option};
    use access_sui::events;

    /// 代表创作者发布的内容。
    /// 这是一个具有唯一ID的对象，由创作者拥有。
    public struct Content has key, store {
        id: UID,
        /// 内容创作者的地址。
        creator: address,
        /// 内容的URI（统一资源标识符）。
        uri: vector<u8>,
        /// 加密后内容的SHA256哈希值。
        content_hash: vector<u8>,
        /// 内容价格，以基础单位计算（例如，SUI的MIST）。
        price: u64,
        /// 存储不同访问级别的动态表（例如，“basic” -> 1, “premium” -> 2）。
        access_levels: Table<vector<u8>, u8>,
    }

    /// 发布新内容。
    public entry fun publish_content(
        creator: address,
        _uri: vector<u8>,
        content_hash: vector<u8>,
        price: u64,
        ctx: &mut TxContext
    ) {
        // 创建一个新的Content对象。
        let mut content = Content {
            id: object::new(ctx),
            creator: creator,
            uri: _uri,
            content_hash: content_hash,
            price: price,
            access_levels: table::new(ctx),
        };

        // 添加一个默认的基础访问级别。
        table::add(&mut content.access_levels, b"basic", 1);

        // 发出内容发布事件。
        events::emit_content_published(object::id(&content), creator, _uri, price);

        // 将Content对象转移给创作者。
        transfer::transfer(content, creator);
    }

    /// 向内容对象添加访问级别。
    public entry fun add_access_level(
        content: &mut Content,
        level_name: vector<u8>,
        level_value: u8,
        ctx: &mut TxContext
    ) {
        // 只有创作者才能添加访问级别。
        assert!(tx_context::sender(ctx) == content.creator, 0);
        table::add(&mut content.access_levels, level_name, level_value);
    }

    /// 从内容对象获取访问级别。
    public fun get_access_level(content: &Content, level_name: vector<u8>): Option<u8> {
        if (table::contains(&content.access_levels, level_name)) {
            option::some<u8>(*table::borrow(&content.access_levels, level_name))
        } else {
            option::none<u8>()
        }
    }

    /// 获取给定内容对象的内容哈希。
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