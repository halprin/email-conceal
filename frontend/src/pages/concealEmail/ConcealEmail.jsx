import { useSelector, useDispatch } from 'react-redux';
import {createConcealedEmail, deleteConcealedEmail, updateConcealedEmail} from './logic';

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
                    <button type='button' className='btn btn-primary' onClick={(event) => dispatch(createConcealedEmail(event.target.form.actualEmailAddress.value, event.target.form.concealEmailDescriptionForAdd.value))}>Create</button>
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
                    <button type='button' className='btn btn-primary' onClick={(event) => dispatch(updateConcealedEmail(event.target.form.concealEmailForUpdate.value, event.target.form.concealEmailDescriptionForUpdate.value))}>Update</button>
                </form>
            </div>

            <div className='row'>
                <h2>Delete</h2>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailForDelete' className='form-label'>Conceal E-mail Address</label>
                        <input type='email' className='form-control' id='concealEmailForDelete' />
                    </div>
                    <button type='button' className='btn btn-primary' onClick={(event) => dispatch(deleteConcealedEmail(event.target.form.concealEmailForDelete.value))}>Delete</button>
                </form>
            </div>
        </div>
    );
}

export default ConcealEmail;
