import { configureStore } from '@reduxjs/toolkit';
import concealedEmailReducer from './pages/concealEmail/concealEmailSlice';
import actualEmailReducer from './pages/actualEmail/actualEmailSlice';

export default configureStore({
    reducer: {
        concealedEmail: concealedEmailReducer,
        actualEmail: actualEmailReducer,
    },
});
