import { setActualEmailAddress } from './actualEmailSlice';
import axios from 'axios';
import { callBackend } from '../../helpers/backend'

export const createActualEmail = (actualEmailAddress) => {
    return async (dispatch, getState) => {
        console.log('Create actual e-mail');
        await createActualEmailInBackend(actualEmailAddress);
        dispatch(setActualEmailAddress(actualEmailAddress));
        console.log('Done Create');
    };
};

const createActualEmailInBackend = async (actualEmailAddress) => {
    await callBackend('/v1/actualEmail', 'post', {
        email: actualEmailAddress,
    });
};
