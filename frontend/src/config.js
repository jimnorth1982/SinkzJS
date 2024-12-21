const serverUrl = 'http://localhost:8080';
const itemsPath = '/items';

const itemsEndpoint = () => `${serverUrl}${itemsPath}`;

export { itemsEndpoint };