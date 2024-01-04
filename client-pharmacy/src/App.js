import React from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import List from './items/List';
import ViewItem from './items/ViewItem'; 
import StickyNavbar from './Navbar';
import GetStarted from './GetStarted';
import InvoiceCreation from './invoices/InvoiceCreation'
import EachInvoice from './invoices/EachInvoice';


export default function App() {

  return (
    <Router>
      <StickyNavbar/>
      
      <Routes>
         <Route path="/items" element={<List />} />
        <Route index element={<  GetStarted />} /> 
         {/* <Route path="/invoices" element={<EachInvoice />} /> */}
            <Route path="/invoices" element={<InvoiceCreation />} />
        <Route path="/product/:id" element={<ViewItem />} />
      </Routes> 
    </Router>
  );
}