import React, { useEffect, useState } from 'react';
import { studentApi } from '../api/studentApi';
import EditStudent from './EditStudent';
import './StudentList.css';
import Pagination from './Pagination'; // 新增导入

function StudentList() {
    const [students, setStudents] = useState([]);
    const [error, setError] = useState('');
    const [editingStudentId, setEditingStudentId] = useState(null);
    const [currentPage, setCurrentPage] = useState(1);
    const studentsPerPage = 10;

    useEffect(() => {
        fetchStudents();
    }, []);

    const fetchStudents = async () => {
        try {
            const response = await studentApi.getAllStudents();
            setStudents(response.data);
        } catch (err) {
            setError('无法获取学生数据');
            console.error(err);
        }
    };

    const handleDelete = async (id) => {
        try {
            await studentApi.deleteStudent(id);
            fetchStudents();
        } catch (err) {
            setError('删除学生失败');
            console.error(err);
        }
    };

    const handleEdit = (id) => {
        setEditingStudentId(id);
    };

    const handleStudentUpdated = () => {
        setEditingStudentId(null);
        fetchStudents();
    };

    const handleCancelEdit = () => {
        setEditingStudentId(null);
    };

    /* 计算当前页的学生 */
    const indexOfLastStudent = currentPage * studentsPerPage;
    const indexOfFirstStudent = indexOfLastStudent - studentsPerPage;
    const currentStudents = students.slice(indexOfFirstStudent, indexOfLastStudent);

    const paginate = (pageNumber) => setCurrentPage(pageNumber);

    return (
        <div className="student-list-container fade-in-up" id="student-list">
            <h2 className="list-title">学生列表</h2>
            {error && <p className="error-message">{error}</p>}
            {editingStudentId ? (
                <EditStudent
                    studentId={editingStudentId}
                    onStudentUpdated={handleStudentUpdated}
                    onCancel={handleCancelEdit}
                />
            ) : (
                <div className="table-container">
                    <table className="student-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>姓名</th>
                                <th>年龄</th>
                                <th>性别</th>
                                <th>邮箱</th>
                                <th>电话</th>
                                <th>操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {currentStudents.map(student => (
                                <tr key={student.id} className="student-row">
                                    <td data-label="ID">{student.id}</td>
                                    <td data-label="姓名">{student.name}</td>
                                    <td data-label="年龄">{student.age}</td>
                                    <td data-label="性别">
                                        <span className={`gender-badge ${student.gender === '男' ? 'male' : 'female'}`}>
                                            {student.gender}
                                        </span>
                                    </td>
                                    <td data-label="邮箱">{student.email}</td>
                                    <td data-label="电话">{student.phone_number}</td>
                                    <td data-label="操作" className="action-buttons">
                                        <button className="edit-btn" onClick={() => handleEdit(student.id)}>
                                            编辑
                                        </button>
                                        <button className="delete-btn" onClick={() => handleDelete(student.id)}>
                                            删除
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                    <Pagination
                        studentsPerPage={studentsPerPage}
                        totalStudents={students.length}
                        paginate={paginate}
                        currentPage={currentPage}
                    />
                </div>
            )}
        </div>
    );
}

export default StudentList;
