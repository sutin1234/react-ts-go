"use client"

import {useState} from "react";
import {buttonStyle, debugStyle, mainStyle} from "@/app/styleProps";

export default function Home() {


    const [foo, setFoo] = useState('')
    async function  getFoo(){
        const r = await fetch('http://localhost:8000/api/v1/getTime')
        const json = await r.json()
        setFoo(json['time'])
    }
  return (
    <main className={mainStyle}>
      <div>
        <div className={debugStyle}>
          <div>
            ENV: {process.env.NODE_ENV}
        </div>
          <div>
            API CONTEXT: {process.env.BASE_URL}:{process.env.PORT}/{process.env.API_FREFIX}/{process.env.API_VERSION}
          </div>

        </div>
        <p>
          foo: {foo}
        </p>
          <button className={buttonStyle} onClick={getFoo}>get foo</button>
      </div>

    </main>
  )
}
