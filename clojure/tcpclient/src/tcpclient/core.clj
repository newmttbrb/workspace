(ns tcpclient.core)

(require '[clojure.java.io :as io])
(require '[clojure.core.async :refer [chan sliding-buffer go go-loop timeout >! <!]])
(import '[java.io StringWriter] '[java.net Socket])

; (defn foo
;   "I don't do a whole lot."
;   [x]
;   (println x "Hello, World!"))

; (defn send-request
;   "Sends an HTTP GET request to the specified host, port, and path"
;   [host port path]
;   (with-open [sock (Socket. host port)
;               writer (io/writer sock)
;               reader (io/reader sock)
;               response (StringWriter.)]
;     (.append writer (str "GET " path "\n"))
;     (.flush writer)
;     (io/copy reader response)
;     (str response)))

; (defn sample
;   []
;   (send-request "google.com" 80 "/"))


(def jump-right-then-left "383;<?xml version=\"1.0\"?>
 <commands>
   <enableAutomation autoFormat=\"true\" cpus=\"0\"/>
   <pressButton cpuid=\"0\" button=\"11\" duration=\"200\"/>
   <pause duration=\"200\"/>
   <pressButton cpuid=\"0\" button=\"10\" duration=\"200\"/>
   <pause duration=\"200\"/>
   <pressButton cpuid=\"0\" button=\"22\" duration=\"500\"/>
   <pause duration=\"200\"/>
   <disableAutomation/>
   <returnQueuedMessages/>
 </commands>")

(def make-a-call "1238;<?xml version=\"1.0\"?>
<commands>
  <enableAutomation autoFormat=\"true\" cpus=\"0 1\"/>
  <pressButton cpuid=\"0\" button=\"12\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"0\" button=\"21\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"0\" button=\"22\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"1\" button=\"37\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"1\" button=\"0\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"0\" button=\"28\" duration=\"200\"/>
  <pause duration=\"300\"/>
  <pressButton cpuid=\"0\" button=\"25\" duration=\"200\"/>
  <pause duration=\"300\"/>
  <pressButton cpuid=\"0\" button=\"25\" duration=\"200\"/>
  <pause duration=\"300\"/>
  <pressButton cpuid=\"0\" button=\"27\" duration=\"200\"/>
  <pause duration=\"300\"/>
  <pressButton cpuid=\"0\" button=\"36\" duration=\"200\"/>
  <pause duration=\"300\"/>
  <pause duration=\"2000\"/>
  <pressButton cpuid=\"0\" button=\"11\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pressButton cpuid=\"1\" button=\"1\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <pause duration=\"10000\"/>
  <pressButton cpuid=\"0\" button=\"21\" duration=\"200\"/>
  <pause duration=\"800\"/>
  <disableAutomation/>
  <returnQueuedMessages/>
</commands>")

 (defn turret-request
    [host port str-to-send]
    (with-open [sock (Socket. host port)
                writer (io/writer sock)
                reader (io/reader sock)
                response (StringWriter.)]
      (.append writer (str str-to-send "\n"))
      (.flush writer)
      (io/copy reader response)
      (println (str response))))

(defn turret-request-async
  [host port]
  (let [in (chan (sliding-buffer 10))]
    (go-loop [data (<! in)]
          (case (first data)
            :exit :exit
            :send (do 
              (turret-request host port (second data)) 
              (recur (<! in)))))
    in))

;;;;;;;;;;;;;;;;;;;;;
;;; async version ;;;
;;;;;;;;;;;;;;;;;;;;;

; (defn turret-request-async-sender
;    [open-socket str-to-send]
;    (with-open [writer (io/writer open-socket)
;                reader (io/reader open-socket)
;                response (StringWriter.)]
;      (.append writer (str str-to-send "\n"))
;      (.flush writer)
;      (io/copy reader response)
;      (println(str response))))

; (defn turret-request-async2
;   [host port]
;   (let [in (chan (sliding-buffer 10))]
;     (go
;       (with-open [sock (Socket. host port)]
;         (loop [data (<! in)]
;           (case (first data)
;             :exit :exit
;             :send (do 
;               (turret-request-async-sender sock (second data)) 
;               (recur (<! in)))))))
;     in))

(defn test-send [ch msg] (go (>! ch msg)))

(def turret1-tcp (turret-request-async "10.204.45.168" 8781))

(defn test-message1 [] (test-send turret1-tcp [:send jump-right-then-left]))
(defn test-message2 [] (test-send turret1-tcp [:send make-a-call]))
(defn test-message3 [] (test-send turret1-tcp [:exit]))
