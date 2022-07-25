import { setConcealedEmailAddress, setConcealedEmailDescription } from './concealEmailSlice';
import { extractUserFromEmailAddress } from '../../helpers/email';
import { callBackend } from '../../helpers/backend'

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
        const concealEmailId = extractUserFromEmailAddress(concealedEmailAddress);
        await updateConcealEmailInBackend(concealEmailId, description);
        dispatch(setConcealedEmailDescription(description));
        console.log('Done Update');
    };
};

export const deleteConcealedEmail = (concealedEmailAddress) => {
    return async (dispatch, getState) => {
        console.log('Delete concealed e-mail');
        const concealEmailId = extractUserFromEmailAddress(concealedEmailAddress);
        await deleteConcealEmailInBackend(concealEmailId);
        dispatch(setConcealedEmailAddress(''));
        dispatch(setConcealedEmailDescription(''));
        console.log('Done Delete');
    };
};

const createConcealEmailInBackend = async (actualEmailAddress, description) => {
    return (await callBackend('/v1/concealEmail', 'post', {
        email: actualEmailAddress,
        description,
    })).concealedEmail;
};

const updateConcealEmailInBackend = async (concealedEmailAddress, description) => {
    await callBackend(`/v1/concealEmail/${concealedEmailAddress}`, 'put', {
        description,
    });
};

const deleteConcealEmailInBackend = async (concealedEmailAddress) => {
    await callBackend(`/v1/concealEmail/${concealedEmailAddress}`, 'delete');
};
