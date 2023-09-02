"use client"

import {useEffect, useState} from "react";
import {buttonStyle, debugStyle, mainStyle} from "@/app/styleProps";

interface TIME_SERVER {
    time:string
}

export default function Home() {
    const [time, setTime] = useState('')
    const [url, setUrl] = useState('');

    const getServerTime = async <T,>(): Promise<T> => {
        const r = await fetch(`${url}/getTime`)
        return await r.json() satisfies T
    }
    const getFoo = async () =>{
        const serverTimeData = await getServerTime<TIME_SERVER>()
        setTime(serverTimeData.time)
    }

    useEffect( ()=>{
        setUrl(`${process.env.BASE_URL}:${process.env.PORT}/${process.env.API_FREFIX}/${process.env.API_VERSION}`)
    },[setUrl])

  return (
    <main className={mainStyle}>
      <div>
        <div className={debugStyle}>
          <div>
            API SERVER TIME: {url}/getTime
          </div>

        </div>
        <p>
          Server Time: {time}
        </p>
          <button className={buttonStyle} onClick={getFoo}>get foo</button>
      </div>

    </main>
  )
}
