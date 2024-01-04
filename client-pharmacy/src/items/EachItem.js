import { Link } from "react-router-dom";
import { API_URL } from "../config";
import axios from "axios"
import React, { useState } from 'react'

export default function EachItem({ item, fetchData }) {
    const [nameValue, setNameValue] = useState(item.name);
    const [categoryValue, setCategoryValue] = useState(item.category);
    const [unitpriceValue, setUnitPriceValue] = useState(item.unitprice);

    const openModal = () => {
        document.getElementById('new-modal-' + item.id).classList.remove("hidden");
    }
    const closeModal = () => {
        document.getElementById('new-modal-' + item.id).classList.add("hidden");
    }
    const completeForm = () => {
        closeModal()
        fetchData()
    }

    const updateItem = (e) => {
        e.preventDefault()
        var form = document.getElementById(`editform-${item.id}`);
        var formData = new FormData(form);
        axios.patch(`${API_URL}/items/${item.id}`, formData)
            .then(res => completeForm())
            .catch(error => console.log(error.response))
    }

    const deleteItem = () => {
        if (window.confirm("Are you sure you want to delete this item??") === true) {
            axios.delete(`${API_URL}/items/${item.id}`)
                .then(res => fetchData())
                .catch(error => console.log(error.response))
        } else {
           console.log("You canceled!");
        }
    }

    return (
        <div className="bg-slate-100 rounded-lg mb-4 p-4 hover:border hover:border-purple-700">
            <div>
                <div>
                    <div className="font-medium">{item.name}</div>
                    <div className="text-slate-400">{item.category}</div>
                    <div className="text-slate-400">{"Rs."+""+item.unitprice}</div>
                </div>
                <div className="text-sm flex space-x-4 mt-4">
                    <Link to={`/product/${item.id}`}>View product</Link>
                    <button onClick={openModal} >Edit</button>
                    <button onClick={deleteItem} className="text-red-600">Delete</button>
                </div>
            </div>

            {/* Start Modal */}
            <div className="relative z-10 hidden" aria-labelledby="modal-title" role="dialog" aria-modal="true" id={`new-modal-${item.id}`}>
                <div className="fixed inset-0 bg-black bg-opacity-70 transition-opacity"></div>

                <div className="fixed z-10 inset-0 overflow-y-auto">
                    <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                        <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

                        <div className="relative inline-block align-middle bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:max-w-lg w-full">
                            <form id={`editform-${item.id}`} onSubmit={updateItem} method="post">
                                <div className="bg-white">
                                    <div className="flex justify-between px-8 py-4 border-b">
                                        <h1 className="font-medium">Update item</h1>
                                        <button type="button" onClick={closeModal}>Close</button>
                                    </div>
                                    <div className="px-8 py-8">
                                        <div className="mb-5">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">Name</label>
                                            <input type="text" name="name" value={nameValue} onChange={(e) => setNameValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                    
                    
                                        <div className="mb-5">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">Category</label>
                                            <input type="text" name="category" value={categoryValue} onChange={(e) => setCategoryValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                        <div className="mb-10">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">UnitPrice</label>
                                            <input type="number" name="unitprice" value={unitpriceValue} onChange={(e) => setUnitPriceValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                        <div className="flex justify-end">
                                            <button className="bg-blue-500 text-white py-1.5 px-4 rounded" type="submit">Submit</button>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
            {/* End Modal */}
        </div>
    )
}