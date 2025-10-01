import { useEffect, useState } from 'react'
import axios from 'axios'

export default function Matters(){
  const [items,setItems]=useState<any[]>([])
  const token = localStorage.getItem('token')||''
  useEffect(()=>{
    axios.get('/api/matters',{headers:{Authorization:`Bearer ${token}`}})
      .then(r=>setItems(r.data))
  },[])
  return (
    <div className="p-6">
      <h1 className="text-xl font-bold mb-4">Matters</h1>
      <Upload />
      <ul className="mt-6 space-y-2">
        {items.map(m=><li key={m.id} className="border p-3 rounded">{m.title} — {m.status}</li>)}
      </ul>
    </div>
  )
}

function Upload(){
  const [title,setTitle]=useState('Acme MSA v1')
  const [text,setText]=useState('Limitation of Liability ...')
  const token = localStorage.getItem('token')||''
  return (
    <div className="border p-3 rounded">
      <h2 className="font-semibold mb-2">Upload & Review</h2>
      <input className="border p-2 w-full mb-2" value={title} onChange={e=>setTitle(e.target.value)} />
      <textarea className="border p-2 w-full h-40" value={text} onChange={e=>setText(e.target.value)} />
      <button className="mt-2 bg-black text-white px-4 py-2" onClick={async ()=>{
        const r = await axios.post('/api/documents',{title, text},{headers:{Authorization:`Bearer ${token}`}})
        const docId = r.data.id
        await axios.post(`/api/review/${docId}/start?playbook_id=1`,{},{
          headers:{Authorization:`Bearer ${token}`}
        })
        location.href=`/review/${docId}`
      }}>Run Review</button>
    </div>
  )
}
