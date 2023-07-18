import React, {useState} from 'react';
import './styles.css';
import {Button, Modal} from "react-bootstrap"; // 引入样式文件

const DataList = ({ data }) => {
    const [showModal1, setShowModal1] = useState(false);
    const [showModal2, setShowModal2] = useState(false);
    const [modalContent1, setModalContent1] = useState('');
    const [modalContent2, setModalContent2] = useState('');

    const handleShowYamlModal = (item) => {
        const params = {
            namespace: item.namespace,
            job_name: item.job_name,
        };

        const queryString = Object.keys(params)
            .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
            .join('&');

        fetch(`/api/v1/operator/job/yaml?${queryString}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                // 处理获取到的数据
                console.log("success resp:", data);
                setModalContent2(data)
                setShowModal2(true);

            })
            .catch(error => {
                // 处理请求错误
                console.error(error);
                setShowModal2(true);
            });
    };
    const handleCloseYamlModal = () => {
        setShowModal2(false);
    };

    const handleShowModal = (item) => {
        const params = {
            namespace: item.namespace,
            job_name: item.job_name,
        };

        const queryString = Object.keys(params)
            .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
            .join('&');

        fetch(`/api/v1/operator/job/log?${queryString}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                // 处理获取到的数据
                console.log("success resp:", data);
                setModalContent1(data)
                setShowModal1(true);

            })
            .catch(error => {
                // 处理请求错误
                console.error(error);
                setShowModal1(true);
            });
    };
    const handleCloseModal = () => {
        setShowModal1(false);
    };

    const deletePod = (record) => {
        fetch(`/api/v1/operator/job/delete`, {
            method: 'POST',
            body: JSON.stringify({
                namespace: record.namespace,
                job_name: record.job_name,
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                // 处理获取到的数据
                console.log("success resp:", data);
                alert('删除成功！');
                window.location.reload();
            })
            .catch(error => {
                // 处理请求错误
                console.error(error);
                alert('删除失败！');
            });
    };

    return (
        <table className="table table-bordered table-striped">
            <thead>
            <tr>
                <th>Job名称</th>
                <th>job状态</th>
                <th>job是否完成</th>
                <th>job开始时间</th>
                <th>job结束时间</th>
                <th>job信息</th>
                <th>yaml</th>
                <th>日志</th>
                <th>删除</th>
            </tr>
            </thead>
            <tbody>
            {data.map((item) => (
                <tr key={item.job_name}>
                    <td>{item.job_name}</td>
                    <td>{item.status}</td>
                    <td className={item.succeeded === 1 ? 'table-success' : 'table-danger'}>
                        {item.succeeded === 1 ? '完成' : '未完成'}
                    </td>
                    <td>{item.startTime}</td>
                    <td>{item.completionTime}</td>
                    <td>{item.message}</td>
                    <td>
                        <Button size="xs" onClick={() => handleShowYamlModal(item)}>View</Button>
                    </td>
                    <td>
                        <Button size="xs" onClick={() => handleShowModal(item)}>View</Button>
                    </td>
                    <td>
                        <Button variant="danger" size="xs" onClick={() => deletePod(item)}>Delete</Button>
                    </td>
                </tr>
            ))}
            </tbody>
            <Modal size="lg" show={showModal1} onHide={handleCloseModal}>
                <Modal.Header closeButton>
                    <Modal.Title>pod日志</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {modalContent1.message}
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleCloseModal}>
                        关闭
                    </Button>
                </Modal.Footer>
            </Modal>
            <Modal show={showModal2} onHide={handleCloseYamlModal}>
                <Modal.Header closeButton>
                    <Modal.Title>yaml内容</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <pre style={{whiteSpace: "pre-wrap"}}>{modalContent2.message}</pre>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleCloseYamlModal}>
                        关闭
                    </Button>
                </Modal.Footer>
            </Modal>
        </table>
    );
};

export default DataList;