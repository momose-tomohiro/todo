FROM node

WORKDIR /go/src/todo/front

RUN npm install --grobal yarn
RUN yarn add @nuxtjs/vuetify
RUN yarn add @nuxtjs/axios
RUN yarn add @nuxtjs/proxy

EXPOSE 3000

CMD ["yarn","dev"]
