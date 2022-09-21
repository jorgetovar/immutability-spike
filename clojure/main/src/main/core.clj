(ns main.core
  (:require [clojure.string :as s])
  (:gen-class)
  )

(defn get-pages-with-code [book]
  (str "The " (s/upper-case (:title book)) " book has " (:pages book) " pages")
  )

(defn change-author [book author]
  (assoc book :author author))


(defn -main
  [& args]
  (let [book {:author "Rich Hickey", :title "Clojure for the Brave & True", :pages 232}
        new-book (change-author book "Daniel Higginbotham")]
    (println "1)" book)
    (println "2)" new-book)
    (println "3)" (get-pages-with-code book))
    )

  )
