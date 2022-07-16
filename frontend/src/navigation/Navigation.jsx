const Navigation = () => {
    return <nav className='navbar navbar-expand-lg navbar-dark bg-dark'>
        <a className='navbar-brand active'>E-mail Conceal</a>
        <button className='navbar-toggler' type='button' data-bs-toggle='collapse'
                data-bs-target='#navbarSupportedContent' aria-controls='navbarSupportedContent' aria-expanded='false'
                aria-label='Toggle navigation'>
            <span className='navbar-toggler-icon'></span>
        </button>
        <div className='collapse navbar-collapse' id='navbarSupportedContent'>
            <div className='navbar-nav'>
                <a className='nav-link' href='./login' target='_blank'>Login</a>
            </div>
        </div>
    </nav>;
}

export default Navigation;