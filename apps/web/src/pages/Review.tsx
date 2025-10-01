import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import axios from 'axios'

export default function Review(){
  const { docId } = useParams()
  const token = localStorage.getItem('token')||''
  const [items,setItems]=useState<any[]>([])
  useEffect(()=>{
    const t = setInterval(()=>{
      axios.get(`/api/review/${docId}/findings`,{headers:{Authorization:`Bearer ${token}`}})
        .then(r=>setItems(r.data))
    }, 1500)
    return ()=>clearInterval(t)
  },[docId])
  return (
    <div className="p-6">
      <h1 className="text-xl font-bold mb-4">Findings</h1>
      <div className="grid gap-3">
        {items.map(f=><FindingCard key={f.id} f={f} />)}
      </div>
    </div>
  )
}

function FindingCard({f}:{f:any}){
  return (
    <div className="border p-3 rounded">
      <div className="text-sm opacity-70">{f.ruleName} — <b>{f.severity}</b></div>
      <div className="mt-2">
        <div className="font-semibold">Clause</div>
        <pre className="whitespace-pre-wrap text-sm">{f.clauseText}</pre>
      </div>
      <div className="mt-2">
        <div className="font-semibold">Rationale</div>
        <p className="text-sm">{f.rationale}</p>
      </div>
      <div className="mt-2">
        <div className="font-semibold">Suggestion</div>
        <pre className="whitespace-pre-wrap text-sm">{f.suggestion||'Pending...'}</pre>
      </div>
    </div>
  )
}
