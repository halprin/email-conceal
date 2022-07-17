const Login = () => {
    return (
        <div className='container'>
            <div className='row'>
                <h1>Login</h1>
            </div>

            <div className='row'>
                <form>
                    <div className='mb-3'>
                        <label htmlFor='username' className='form-label'>E-mail Address</label>
                        <input type='email' className='form-control' id='username' />
                    </div>
                    <div className='mb-3'>
                        <label htmlFor='password' className='form-label'>Password</label>
                        <input type='password' className='form-control' id='password' />
                    </div>
                    <button type='submit' className='btn btn-primary'>Login</button>
                </form>
            </div>
        </div>
    );
}

export default Login;
