(ns async.core)

(require '[clojure.core.async :refer [chan sliding-buffer go go-loop timeout >! <!]])

(defn database-consumer
  []
  (let [in (chan (sliding-buffer 64))]
    (go-loop [data (<! in)]
      (when data
        (println (format "database consumer received data %s" data))
        (recur (<! in))))
    in))

(defn sse-consumer
  []
  (let [in (chan (sliding-buffer 64))]
    (go-loop [data (<! in)]
      (when data
        (println (format "sse consumer received data %s" data))
        (recur (<! in))))
    in))

(defn messages
  []
  (range 4))

(defn producer
  [& channels]
  (go
    (doseq [msg (messages)
            out channels]
      (<! (timeout 100))
      (>! out msg))))

(producer (database-consumer) (sse-consumer)) 
