import {Link} from 'react-router-dom';

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
                        <Link className='nav-link' to='./actual-email'>Actual E-mails</Link>
                    </li>
                    <li className='nav-item'>
                        <Link className='nav-link' to='./conceal-email'>Concealed E-mails</Link>
                    </li>
                    <li className='nav-item'>
                        <Link className='nav-link' to='./login'>Login</Link>
                    </li>
                </ul>
            </div>
        </nav>
    );
}

export default Navigation;