import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    address: null,
};

const actualEmailSlice = createSlice({
    name: 'actualEmail',
    initialState,
    reducers: {
        setActualEmailAddress: (state, action) => {
            state.address = action.payload;
        },
    },
});

export const { setActualEmailAddress } = actualEmailSlice.actions

export default actualEmailSlice.reducer
