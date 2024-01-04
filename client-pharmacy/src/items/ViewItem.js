import React, { useState, useEffect } from 'react'
import { API_URL } from '../config';
import { useParams, Link } from "react-router-dom";


export default function ViewItem() {
  let { id } = useParams();
  const [item, setItem] = useState({})
  const [loading, setLoading] = useState(false)

  

  useEffect(() => {
    const fetchItem = async () => {
      setLoading(true)
    //     // Log the current value of id
     console.log("Current ID:", id);

    // // Handle the case when id is not available
    if (!id) {
     setLoading(false);
      return;
  }
      try {
        const response = await fetch(`${API_URL}/items/${id}`);
        const json = await response.json();
        setItem(json.data);
        setLoading(false)
      } catch (error) {
        console.log("error", error);
        setLoading(false);
      }
    };

    fetchItem();
  }, [id]);

  return (
    <div>
      {!loading ?
        <div className="flex justify-center">
          <div className="lg:w-1/3 w-full">
            <div className="p-10">
              <div className="mb-10 flex items-center justify-between">
                <Link to="/items"><h1 className="font-bold">Go back</h1></Link>
              </div>
              <div className="bg-slate-200 rounded-lg px-5">
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Name</div>
                  <div className="text-slate-800 font-medium">{item.name}</div>
                </div>
               
                <div className="flex border-b py-4">
                  <div className="mr-4 text-slate-400">Category</div>
                  <div className="text-slate-800 font-medium">{item.category}</div>
                </div>
                <div className="flex py-4">
                  <div className="mr-4 text-slate-400">UnitPrice(LKR)</div>
                  <div className="text-slate-800 font-medium">{item.unitprice}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        : ''}
    </div>
  )
}