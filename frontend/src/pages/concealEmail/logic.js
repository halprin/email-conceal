import { setConcealedEmailAddress, setConcealedEmailDescription } from './concealEmailSlice';

export const createConcealedEmail = async () => {
    console.log('Create');
    setConcealedEmailAddress('george@example.com');
    console.log('Done Create');
};
