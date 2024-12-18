import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import AddStudent from './components/AddStudent';
import StudentList from './components/StudentList';
import Navbar from './components/Navbar';
import Home from './components/Home';
import SystemOverview from './components/SystemOverview';

function App() {
    const handleStudentAdded = () => {
        // 可以添加刷新列表或其他需要的操作
        console.log('学生添加成功，可以在这里刷新列表');
    };

    return (
        <Router>
            <div className="App">
                <Navbar />
                <main className="App-main fade-in-up">
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/student-list" element={<StudentList />} />
                        <Route path="/add-student" element={<AddStudent onStudentAdded={handleStudentAdded} />} />
                        <Route path="/system-overview" element={<SystemOverview />} />
                    </Routes>
                </main>
            </div>
        </Router>
    );
}

export default App;
