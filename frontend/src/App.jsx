import Navigation from './components/navigation/Navigation';
import Welcome from './pages/Welcome';
import { useOutlet } from 'react-router-dom';

const App = () => {
    const outlet = useOutlet();
    const welcome = <Welcome />;
    return (
        <div className='container'>
            <Navigation />
            {outlet || welcome}
        </div>
    );
}

export default App;
