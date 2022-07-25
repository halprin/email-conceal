import { useSelector, useDispatch } from 'react-redux';
import {createActualEmail} from './logic';

const ActualEmail = () => {
    const dispatch = useDispatch();
    const actualEmailAddress = useSelector(state => state.actualEmail.address);

    return (
        <div className='container'>
            <div className='row'>
                <h1>Actual E-mails</h1>
            </div>

            <div className='row'>
                <h2>Actual E-mail State</h2>
                <h3>E-mail Address:</h3>
                <span>{actualEmailAddress}</span>
            </div>

            <div className='row'>
                <h2>Create</h2>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='actualEmailAddress' className='form-label'>Real E-mail Address</label>
                        <input type='email' className='form-control' id='actualEmailAddress' />
                    </div>
                    <button type='button' className='btn btn-primary' onClick={(event) => dispatch(createActualEmail(event.target.form.actualEmailAddress.value))}>Create</button>
                </form>
            </div>
        </div>
    );
}

export default ActualEmail;
