import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    address: null,
    description: null,
};

const concealedEmailSlice = createSlice({
    name: 'concealedEmail',
    initialState,
    reducers: {
        setConcealedEmailAddress: (state, action) => {
            state.address = action.payload;
        },
        setConcealedEmailDescription: (state, action) => {
            state.description = action.payload;
        },
    },
});

export const { setConcealedEmailAddress, setConcealedEmailDescription } = concealedEmailSlice.actions

export default concealedEmailSlice.reducer
