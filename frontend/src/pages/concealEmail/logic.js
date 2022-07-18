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

const createConcealEmailInBackend = async (actualEmailAddress, description) => {
    const response = await axios.post(`http://localhost:8000/v1/concealEmail`, {
        email: actualEmailAddress,
        description,
    });

    return response.data.concealedEmail;
};
