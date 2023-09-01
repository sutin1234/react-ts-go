"use client"

import {css} from '@panda/css'
import {useState} from "react";

export default function Home() {
  const mainStyle = css({
    fontSize: "2xl",
    p: 4
  })
  const buttonStyle = css({
    bgColor:"blue.500",
    p:2,
    width: 180,
    color:"white",
    borderRadius: 20
  })
    const [foo, setFoo] = useState('')
    async function  getFoo(){
        const r = await fetch('http://localhost:3001/')
        const json = await r.json()
        setFoo(json['foo'])
    }
  return (
    <main className={mainStyle}>
      <div>
        <p>
          foo: {foo}
        </p>
          <button className={buttonStyle} onClick={getFoo}>get foo</button>
      </div>

    </main>
  )
}
