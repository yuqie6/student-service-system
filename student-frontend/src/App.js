import React, { useState } from 'react';
import './App.css';
import AddStudent from './components/AddStudent';
import StudentList from './components/StudentList';
import Navbar from './components/Navbar'; // 新增导入

function App() {
  const [refresh, setRefresh] = useState(false);

  const handleStudentAdded = () => {
    setRefresh(!refresh);
  };

  return (
    <div className="App">
      <Navbar /> {/* 添加导航栏 */}
      <header className="App-header">
        <div className="header-content">
          <h1>学生信息管理系统</h1>
          <p className="header-subtitle">Student Information Management System</p>
        </div>
      </header>
      <main className="App-main fade-in-up">
        <div className="content-wrapper">
          <section className="add-student-section">
            <AddStudent onStudentAdded={handleStudentAdded} />
          </section>
          <section className="list-student-section">
            <StudentList key={refresh} />
          </section>
        </div>
      </main>
      <footer className="App-footer">
        <p>© 2024 学生管理系统 - 专业版</p>
      </footer>
    </div>
  );
}

export default App;
