import { useState } from 'react'
import axios from 'axios'

export default function Login(){
  const [email,setEmail]=useState('demo@demo.dev')
  const [password,setPassword]=useState('demo')
  return (
    <div className="p-8 max-w-sm mx-auto">
      <h1 className="text-xl font-bold mb-4">Login</h1>
      <input className="border p-2 w-full mb-2" value={email} onChange={e=>setEmail(e.target.value)} />
      <input className="border p-2 w-full mb-2" type="password" value={password} onChange={e=>setPassword(e.target.value)} />
      <button className="bg-black text-white px-4 py-2" onClick={async ()=>{
        const r = await axios.post('/api/auth/login',{email,password})
        localStorage.setItem('token', r.data.token)
        location.href='/matters'
      }}>Sign in</button>
    </div>
  )
}
