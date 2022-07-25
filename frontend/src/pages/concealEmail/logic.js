import { setConcealedEmailAddress, setConcealedEmailDescription } from './concealEmailSlice';
import axios from 'axios';

export const createConcealedEmail = (actualEmailAddress, description) => {
    return async (dispatch, getState) => {
        console.log('Create concealed e-mail');
        const concealedEmailAddress = await createConcealEmailInBackend(actualEmailAddress, description);
        dispatch(setConcealedEmailAddress(concealedEmailAddress));
        dispatch(setConcealedEmailDescription(description));
        console.log('Done Create');
    };
};

export const updateConcealedEmail = (concealedEmailAddress, description) => {
    return async (dispatch, getState) => {
        console.log('Update concealed e-mail');
        await updateConcealEmailInBackend(concealedEmailAddress, description);
        dispatch(setConcealedEmailDescription(description));
        console.log('Done Update');
    };
};

export const deleteConcealedEmail = (concealedEmailAddress) => {
    return async (dispatch, getState) => {
        console.log('Delete concealed e-mail');
        await deleteConcealEmailInBackend(concealedEmailAddress);
        dispatch(setConcealedEmailAddress(''));
        dispatch(setConcealedEmailDescription(''));
        console.log('Done Delete');
    };
};

const updateConcealEmailInBackend = async (concealedEmailAddress, description) => {
    await axios.put(`http://localhost:8000/v1/concealEmail/${concealedEmailAddress}`, {
        description,
    });
};

const createConcealEmailInBackend = async (actualEmailAddress, description) => {
    const response = await axios.post(`http://localhost:8000/v1/concealEmail`, {
        email: actualEmailAddress,
        description,
    });

    return response.data.concealedEmail;
};

const deleteConcealEmailInBackend = async (concealedEmailAddress) => {
    await axios.delete(`http://localhost:8000/v1/concealEmail/${concealedEmailAddress}`);
};
