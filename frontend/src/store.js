import { configureStore } from '@reduxjs/toolkit';
import concealedEmailReducer from './pages/concealEmail/concealEmailSlice';

export default configureStore({
    reducer: {
        concealedEmail: concealedEmailReducer,
    },
});
