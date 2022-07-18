import { setConcealedEmailAddress, setConcealedEmailDescription } from './concealEmailSlice';

export const createConcealedEmail = (actualEmailAddress, description) => {
    console.log('actualEmailAddress', actualEmailAddress);
    console.log('description', description);
    return async (dispatch, getState) => {
        console.log('Create concealed e-mail');
        dispatch(setConcealedEmailAddress('george@example.com'));
        console.log('Done Create');
    };
};

const createConcealEmailInBackend = async (actualEmailAddress, description) => {

};
