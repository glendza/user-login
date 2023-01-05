type AvailableKeys = "token";

const storage = {
  getItem: <T>(key: AvailableKeys): T => {
    return window.localStorage.getItem(key) as T;
  },
  setItem: (key: AvailableKeys, value: unknown): void => {
    window.localStorage.setItem(key, JSON.stringify(value));
  },
  removeItem: (key: AvailableKeys): void => {
    window.localStorage.removeItem(key);
  },
};

export default storage;
