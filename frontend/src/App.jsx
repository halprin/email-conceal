import Navigation from './components/navigation/Navigation';
import { Outlet } from 'react-router-dom';

const App = () => {
    return (
        <div className='container'>
            <Navigation />
            <Outlet />
        </div>
    );
}

export default App;
