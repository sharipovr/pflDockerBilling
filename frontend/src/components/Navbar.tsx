import React from 'react';
import { Link } from 'react-router-dom';

const Navbar: React.FC = () => (
  <nav className="bg-gray-800 p-4 text-white">
    <Link to="/dashboard" className="mr-4">
      Dashboard
    </Link>
    <Link to="/" className="mr-4">
      Выйти
    </Link>
  </nav>
);

export default Navbar;