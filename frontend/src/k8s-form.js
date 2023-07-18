import React, { useState,useEffect } from 'react';
import { Button, Modal } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import DataList from './k8s-data-list';

function MyForm() {
    // 使用useState来定义表单中的多个参数的状态
    const [name, setName] = useState('');
    const [message, setMessage] = useState('');
    const [jobName, setJobName] = useState('');
    const [image, setImage] = useState('');

    const [showModal, setShowModal] = useState(false);
    const [modalContent, setModalContent] = useState('');

    const handleShowModal = () => {
        setShowModal(true);
    };
    const handleCloseModal = () => {
        setShowModal(false);
    };

    const [data, setData] = useState([]);

    const handleSubmit = (e) => {
        e.preventDefault();

        fetch(`/api/v1/operator/job/operator`, {
            method: 'POST',
            body: JSON.stringify({
                namespace: name,
                args: message,
                name: jobName,
                image: image,
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                // 处理获取到的数据
                console.log("success resp:", data);
                // 提交后重置表单参数
                // setName('');
                setMessage('');
                // 处理获取到的数据，设置弹框内容
                setModalContent('提交成功！');
                handleShowModal();

            })
            .catch(error => {
                // 处理请求错误
                console.error(error);
                // 提交后重置表单参数
                // setName('');
                setMessage('');
                setModalContent('提交失败！');
                handleShowModal();
            });
    };

    useEffect(() => {
        const defaultNamespace = 'algorithm';
        setName(defaultNamespace);
        setJobName('demo-job');
        setImage('alpine:latest');
        const fetchData = () => {
            const params = {
                namespace: defaultNamespace,
            };

            const queryString = Object.keys(params)
                .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
                .join('&');

            fetch(`/api/v1/operator/job/list?${queryString}`,{
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            })
                .then(response => response.json())
                .then(data => {
                    if (data.list !== null) {
                        setData(data.list); // 更新数据状态
                    }
                })
                .catch(error => {
                    console.error('请求失败', error);
                    clearInterval(interval);
                });
        };

        // 初始请求数据
        fetchData();

        // 每2秒请求一次数据
        const interval = setInterval(() => {
            fetchData();
        }, 5000);

        // 清除定时器
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="container">
            <h1>随时随地启动job!</h1>
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label htmlFor="name" className="form-label">namespace:</label>
                    <input type="text" className="form-control" id="name"  value={name} onChange={(e) => setName(e.target.value)} />
                </div>
                <div className="mb-3">
                    <label htmlFor="jobName" className="form-label">jobName:</label>
                    <input type="text" className="form-control" id="jobName"  value={jobName} onChange={(e) => setJobName(e.target.value)} />
                </div>
                <div className="mb-3">
                    <label htmlFor="image" className="form-label">image:</label>
                    <input type="text" className="form-control" id="image"  value={image} onChange={(e) => setImage(e.target.value)} />
                </div>
                <div className="mb-3">
                    <label htmlFor="message" className="form-label">参数:</label>
                    <textarea className="form-control" required={true} id="message" rows="4" value={message} onChange={(e) => setMessage(e.target.value)}></textarea>
                </div>
                <button type="submit" className="btn btn-primary">发射！</button>
            </form>
            <Modal show={showModal} onHide={handleCloseModal}>
                <Modal.Header closeButton>
                    <Modal.Title>提交状态</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {modalContent}
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleCloseModal}>
                        关闭
                    </Button>
                </Modal.Footer>
            </Modal>
            <br />
            <br />
            <div>
                <h2>任务列表(每5s刷新)</h2>
            </div>
            {/* 数据列表组件 */}
            <DataList data={data} />
        </div>
    );
}

export default MyForm;