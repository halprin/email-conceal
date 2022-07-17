const Navigation = () => {
    return (
        <nav className='navbar navbar-expand-lg navbar-dark bg-dark'>
            <a className='navbar-brand active' href='/'>E-mail Conceal</a>
            <button className='navbar-toggler' type='button' data-bs-toggle='collapse'
                    data-bs-target='#navbarSupportedContent' aria-controls='navbarSupportedContent' aria-expanded='false'
                    aria-label='Toggle navigation'>
                <span className='navbar-toggler-icon'></span>
            </button>
            <div className='collapse navbar-collapse' id='navbarSupportedContent'>
                <ul className='navbar-nav'>
                    <li className='nav-item'>
                        <a className='nav-link' href='./conceal-email'>Concealed E-mails</a>
                    </li>
                    <li className='nav-item'>
                        <a className='nav-link' href='./login'>Login</a>
                    </li>
                </ul>
            </div>
        </nav>
    );
}

export default Navigation;