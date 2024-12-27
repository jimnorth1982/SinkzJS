/**
 * @typedef {Object} ItemType
 * @property {string} id
 * @property {string} name
 */

/**
 * @typedef {Object} Rarity
 * @property {string} id
 * @property {string} name
 */

/**
 * @typedef {Object} Image
 * @property {string} id
 * @property {string} url
 */

/**
 * @typedef {Object} Attribute
 * @property {string} id
 * @property {string} name
 * @property {number} lowValue
 * @property {number} highValue
 * @property {AttributeGrouping} attributeGrouping
 */

/**
 * @typedef {Object} AttributeGrouping
 * @property {string} id
 * @property {string} name
 */

/**
 * @typedef {Object} Item
 * @property {string} id
 * @property {string} name
 * @property {number} requiredLevel
 */

/**
 * @typedef {Object} ItemsResponse
 * @property {Array<Item>} items
 * @property {string} http_status
 * @property {string} message
 */

/**
 * @typedef {Object} Url
 * @property {string} url - The full URL constructed from the host and path.
 * @property {string} host - The host part of the URL.
 * @property {string} path - The path part of the URL.
 */

/**
 * Represents an exile with basic equipment IDs.
 * @typedef {Object} Exile
 * @property {number} id - The unique identifier of the exile.
 * @property {string} name - The name of the exile.
 * @property {number} level - The level of the exile.
 * @property {number} helmet_id - The ID of the helmet equipped by the exile.
 * @property {number} armor_id - The ID of the armor equipped by the exile.
 * @property {number} weapon_id - The ID of the weapon equipped by the exile.
 * @property {number} shield_id - The ID of the shield equipped by the exile.
 * @property {number} boots_id - The ID of the boots equipped by the exile.
 */

/**
 * Represents an item with basic properties.
 * @typedef {Object} Item
 * @property {number} id - The unique identifier of the item.
 * @property {string} name - The name of the item.
 * @property {string} type - The type of the item.
 * @property {string} rarity - The rarity of the item.
 * @property {string} imageUrl - The URL of the item's image.
 */

/**
 * Represents an exile with fully hydrated equipment items.
 * @typedef {Object} HydratedExile
 * @property {number} id - The unique identifier of the exile.
 * @property {string} name - The name of the exile.
 * @property {number} level - The level of the exile.
 * @property {Item} helmet - The helmet item equipped by the exile.
 * @property {Item} armor - The armor item equipped by the exile.
 * @property {Item} weapon - The weapon item equipped by the exile.
 * @property {Item} shield - The shield item equipped by the exile.
 * @property {Item} boots - The boots item equipped by the exile.
 */

/**
 * Represents a response containing a list of exiles.
 * @typedef {Object} ExilesResponse
 * @property {string} message - The response message.
 * @property {number} http_status - The HTTP status code of the response.
 * @property {Exile[]} exiles - The list of exiles in the response.
 */