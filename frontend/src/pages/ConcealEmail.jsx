const ConcealEmail = () => {
    return (
        <div className='container'>
            <div className='row'>
                <h1>Create</h1>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='actualEmailAddress' className='form-label'>Real E-mail Address</label>
                        <input type='email' className='form-control' id='actualEmailAddress' />
                    </div>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailDescriptionForAdd' className='form-label'>Description</label>
                        <input type='text' className='form-control' id='concealEmailDescriptionForAdd' />
                    </div>
                    <button type='submit' className='btn btn-primary'>Create</button>
                </form>
            </div>

            <div className='row'>
                <h1>Update</h1>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailForUpdate' className='form-label'>Conceal E-mail Address</label>
                        <input type='email' className='form-control' id='concealEmailForUpdate' />
                    </div>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailDescriptionForUpdate' className='form-label'>Description</label>
                        <input type='text' className='form-control' id='concealEmailDescriptionForUpdate' />
                    </div>
                    <button type='submit' className='btn btn-primary'>Update</button>
                </form>
            </div>

            <div className='row'>
                <h1>Delete</h1>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='concealEmailForDelete' className='form-label'>Conceal E-mail Address</label>
                        <input type='email' className='form-control' id='concealEmailForDelete' />
                    </div>
                    <button type='submit' className='btn btn-primary'>Delete</button>
                </form>
            </div>
        </div>
    );
}

export default ConcealEmail;
