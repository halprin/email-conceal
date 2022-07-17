import { useSelector, useDispatch } from 'react-redux';
import { setConcealedEmailAddress, setConcealedEmailDescription } from './concealEmailSlice';
import {createConcealedEmail} from './logic';

const ConcealEmail = () => {
    const dispatch = useDispatch();
    const concealedEmailAddress = useSelector(state => state.concealedEmail.address);
    const concealedEmailDescription = useSelector(state => state.concealedEmail.description);

    return (
        <div className='container'>
            <div className='row'>
                <h1>Concealed E-mails</h1>
            </div>

            <div className='row'>
                <h2>Concealed E-mail State</h2>
                <h3>E-mail Address:</h3>
                <span>{concealedEmailAddress}</span>
                <h3>Description:</h3>
                <span>{concealedEmailDescription}</span>
            </div>

            <div className='row'>
                <h2>Create</h2>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='actualEmailAddress' className='form-label'>Real E-mail Address</label>
                        <input type='email' className='form-control' id='actualEmailAddress' />
                    </div>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailDescriptionForAdd' className='form-label'>Description</label>
                        <input type='text' className='form-control' id='concealEmailDescriptionForAdd' />
                    </div>
                    <button type='button' className='btn btn-primary' onClick={() => dispatch(createConcealedEmail)}>Create</button>
                </form>
            </div>

            <div className='row'>
                <h2>Update</h2>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailForUpdate' className='form-label'>Conceal E-mail Address</label>
                        <input type='email' className='form-control' id='concealEmailForUpdate' />
                    </div>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailDescriptionForUpdate' className='form-label'>Description</label>
                        <input type='text' className='form-control' id='concealEmailDescriptionForUpdate' />
                    </div>
                    <button type='button' className='btn btn-primary'>Update</button>
                </form>
            </div>

            <div className='row'>
                <h2>Delete</h2>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailForDelete' className='form-label'>Conceal E-mail Address</label>
                        <input type='email' className='form-control' id='concealEmailForDelete' />
                    </div>
                    <button type='button' className='btn btn-primary'>Delete</button>
                </form>
            </div>
        </div>
    );
}

export default ConcealEmail;
