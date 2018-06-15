import Axios from 'axios';

export const getRandom = async (): Promise<string[]> => {
    const url = 'http://localhost:3000/get/random';
    const resp = await Axios.get(url);
    const result = resp.data as string[];
    return result;
};
