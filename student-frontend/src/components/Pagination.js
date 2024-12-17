
import React from 'react';
import './Pagination.css';

function Pagination({ studentsPerPage, totalStudents, paginate, currentPage }) {
    const pageNumbers = [];

    for (let i = 1; i <= Math.ceil(totalStudents / studentsPerPage); i++) {
        pageNumbers.push(i);
    }

    return (
        <nav className="pagination-nav">
            <ul className="pagination-list">
                {pageNumbers.map(number => (
                    <li key={number} className={`pagination-item ${currentPage === number ? 'active' : ''}`}>
                        <a onClick={() => paginate(number)} href="#!">
                            {number}
                        </a>
                    </li>
                ))}
            </ul>
        </nav>
    );
}

export default Pagination;