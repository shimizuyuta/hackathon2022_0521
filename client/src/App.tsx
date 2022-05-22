import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Login from './pages/Login/Login';
import Top from './pages/Top/Top';
import SignUp from './pages/SiginUp/SignUp';
import './App.css';
import Navbar from './components/Navbar';
import Post from './pages/Posts/Post';
import Profile from './pages/Profile/Profile';

function App() {
  return (
    <>
    <BrowserRouter>
      <Navbar/>
      <Routes>
        <Route path='/' element ={<Top />} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<SignUp />} />
        <Route path='/post' element={<Post />} />
        <Route path='/profile' element={<Profile />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}

export default App;
