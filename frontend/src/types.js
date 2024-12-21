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