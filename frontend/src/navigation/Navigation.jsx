const Navigation = () => {
    return (
        <nav className='navbar navbar-expand-lg navbar-dark bg-dark'>
            <a className='navbar-brand active'>E-mail Conceal</a>
            <button className='navbar-toggler' type='button' data-bs-toggle='collapse'
                    data-bs-target='#navbarSupportedContent' aria-controls='navbarSupportedContent' aria-expanded='false'
                    aria-label='Toggle navigation'>
                <span className='navbar-toggler-icon'></span>
            </button>
            <div className='collapse navbar-collapse' id='navbarSupportedContent'>
                <ul className='navbar-nav'>
                    <li className='nav-item'>
                        <a className='nav-link' href='./create-conceal-email'>Create Concealed E-mail</a>
                    </li>
                    <li className='nav-item'>
                        <a className='nav-link' href='./update-conceal-email'>Update Concealed E-mail</a>
                    </li>
                    <li className='nav-item'>
                        <a className='nav-link' href='./delete-conceal-email'>Delete Concealed E-mail</a>
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