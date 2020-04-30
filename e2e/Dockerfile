FROM cypress/base:10
COPY . .
#COPY package.json .
RUN npm i
CMD ["npm", "run", "cypress:run"]
