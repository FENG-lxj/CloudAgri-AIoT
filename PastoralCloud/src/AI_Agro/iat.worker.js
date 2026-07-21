// iat.worker.js
self.onmessage = function (e) {
    const { command, data } = e.data
    if (command === 'start') {
        startListening(data)
    } else if (command === 'stop') {
        stopListening()
    }
}

let ws = null
let audioData = []
let audioChunks = []
let isFinal = false

function startListening({ appId, apiKey, apiSecret, language, accent }) {
    const host = 'iat-api.xfyun.cn'
    const url = `wss://${host}/v2/iat`
    const date = new Date().toGMTString()
    const algorithm = 'hmac-sha256'
    const headers = 'host date request-line'
    const signatureOrigin = `host: ${host}\ndate: ${date}\nGET /v2/iat HTTP/1.1`

    const signature = CryptoJS.enc.Base64.stringify(
        CryptoJS.HmacSHA256(signatureOrigin, apiSecret)
    )

    const authorizationOrigin = `api_key="${apiKey}", algorithm="${algorithm}", headers="${headers}", signature="${signature}"`
    const authorization = btoa(authorizationOrigin)

    const finalUrl = `${url}?authorization=${authorization}&date=${date}&host=${host}`

    ws = new WebSocket(finalUrl)

    ws.onopen = () => {
        const params = {
            common: { app_id: appId },
            business: {
                language,
                domain: 'iat',
                accent,
                vad_eos: 5000,
                dwa: 'wpgs'
            },
            data: {
                status: 0,
                format: 'audio/L16;rate=16000',
                encoding: 'raw',
                audio: ''
            }
        }
        ws.send(JSON.stringify(params))
        self.postMessage({ type: 'status', status: 'started' })
    }

    ws.onmessage = (e) => {
        const result = JSON.parse(e.data)
        if (result?.payload?.result) {
            const text = result.payload.result.text
            const isFinal = result.payload.result.is_last === 1
            self.postMessage({ type: 'result', text, isFinal })
        }
    }

    ws.onerror = (e) => {
        self.postMessage({ type: 'error', error: e })
    }

    ws.onclose = () => {
        self.postMessage({ type: 'status', status: 'closed' })
    }

    audioChunks = []
}

function stopListening() {
    if (ws && ws.readyState === WebSocket.OPEN) {
        const finalData = {
            common: { app_id: '' },
            business: {},
            data: {
                status: 2,
                format: 'audio/L16;rate=16000',
                encoding: 'raw',
                audio: ''
            }
        }
        ws.send(JSON.stringify(finalData))
        ws.close()
    }
    audioChunks = []
}

function sendAudioChunk(chunk) {
    if (ws.readyState !== WebSocket.OPEN) return
    const base64 = btoa(String.fromCharCode.apply(null, new Uint8Array(chunk)))
    const params = {
        common: { app_id: '' },
        business: {},
        data: {
            status: 1,
            format: 'audio/L16;rate=16000',
            encoding: 'raw',
            audio: base64
        }
    }
    ws.send(JSON.stringify(params))
}

// 模拟 CryptoJS（在 worker 中无法直接引入，需打包或内联）
// 此处略去，实际使用请通过 webpack 打包引入