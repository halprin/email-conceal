import axios from 'axios';

export const callBackend = async (urlPath, action, requestBody) => {
    const response = await axios({
        method: action,
        url: `http://localhost:8000${urlPath}`,
        data: requestBody,
    });

    return response.data;
};
