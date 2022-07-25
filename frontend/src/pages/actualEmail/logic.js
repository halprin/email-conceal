import { setActualEmailAddress } from './actualEmailSlice';
import axios from 'axios';

export const createActualEmail = (actualEmailAddress) => {
    return async (dispatch, getState) => {
        console.log('Create actual e-mail');
        await createActualEmailInBackend(actualEmailAddress);
        dispatch(setActualEmailAddress(actualEmailAddress));
        console.log('Done Create');
    };
};

const createActualEmailInBackend = async (actualEmailAddress) => {
    await axios.post(`http://localhost:8000/v1/actualEmail`, {
        email: actualEmailAddress,
    });
};
