#FROM node:12
#WORKDIR /root/
#COPY package.json .
#RUN npm install

FROM cypress/base:10
WORKDIR /root

COPY package.json .
COPY package-lock.json .
#COPY cypress.json .
#ENV CI=1
RUN npm i

RUN npx cypress verify

#RUN $(npm bin)/cypress run
CMD ["npm", "run", "cypress:run"]
