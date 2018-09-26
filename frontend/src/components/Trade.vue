<template>
  <div class="trade-page">

    <section id="main-wrapper" :style="{'background-image': `url(${require('../assets/images/bg.jpg')})`}">
      <div class="container-fluid">
        <div class="main-wrap-table">
          <div class="table-cell-left">
            <div class="container-fluid">
              <div class="row">
                <div class="col-lg-12">
                  <div class="main-wrapper-top">
                    <ul>
                      <li v-b-tooltip.hover title="not yet!"><span>2</span> hours</li>
                      <li v-b-tooltip.hover title="not yet!"><span>5</span> hours</li>
                      <li v-b-tooltip.hover title="not yet!"><span>15</span> minutes</li>
                      <li v-b-tooltip.hover title="not yet!"><span>30</span> minutes</li>
                    </ul>
                  </div>

                  <Chart @chartLoaded="chartLoaded" :simpleDataFromTrade="simpleData" :priceFromTrade="currentprice"/>

                </div>
              </div>
            </div>
          </div>

          <div class="table-cell-right">
            <div class="expire-time text-center">
              <p>Expiration Time</p>
              <form action="">
                <select v-model="expiry" class="minimal">
                  <option>1 Minute</option>
                  <option>5 Minutes</option>
                  <option>15 Minutes</option>
                </select>
              </form>
            </div>

            <div class="potential-profit">
              <div class="profit-cell">
                <span>{{percentage}}</span>
              </div>
              <div class="profit-cell-right">
                <ul>
                  <li>Potential Profit</li>
                  <li>DEMO {{potentialProfit}}</li>
                </ul>
              </div>
            </div>

            <div class="market-price text-center">
              <p>Market Price</p>
              <!-- <input v-model="currentprice" type="text"> -->
              <div>{{currentprice}}</div>
            </div>

            <div class="trade-amount text-center">
              <p>Bet Amount</p>
              <input v-model="betsize" id="betsize" class="trade-rate">
              <div class="trade-block">
                <button @click="betsize /= 2">/ 2</button>
                <button @click="betsize *= 2">x 2</button>
              </div>
            </div>
            <b-tooltip v-if="demoBalance < betsize" show target="betsize" placement="bottom">
              You Don't have enough!
            </b-tooltip>

            <div class="buy-sell text-center">
              <button @click="newBet('buy')"><i class="fas fa-chart-line"></i>Buy</button>
              <button @click="newBet('sell')"><i class="fas fa-chart-line"></i>Sell</button>
              <button v-if="demoBalance === 0" @click="getMoreDemo()" style="font-size: 14px;">Get More Demo Token</button>
            </div>
          </div>
        </div>

        <div class="row">
          <div class="col-lg-5 col-md-5">
            <div class="trades-section">
              <h5 class="section-head">
                Trade History -
                <span v-if="expiry === '15 Minutes'" class="section-head-span">{{expiry.substring(0, 4)}}</span>
                <span v-else class="section-head-span">{{expiry.substring(0, 3)}}</span>
              </h5>
              <div class="table-responsive">
                <table class="table">
                  <thead>
                    <tr>
                      <td>Assets</td>
                      <td>Amount</td>
                      <td>Start/Strike Price</td>
                      <td>Result</td>
                      <td>Returns</td>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(doneBet, index) of doneBets" :key="`doneBet-${index}`">
                      <td v-if="doneBet.position === 'buy'"><i class="fas fa-caret-up color-green"></i> {{doneBet.pair}}</td>
                      <td v-else-if="doneBet.position === 'sell'"><i style="transform: rotate(180deg);" class="fas fa-caret-up color-red"></i> {{doneBet.pair}}</td>
                      <td>{{doneBet.betsize}}</td>
                      <td>{{doneBet.price}}/{{doneBet.endPrice}}</td>
                      <td>{{doneBet.result}}</td>
                      <td>{{doneBet.returns}}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
          <div class="col-lg-4 col-md-4">
            <div class="order-pending">
              <h5 class="section-head">
                Pending Orders -
                <span v-if="expiry === '15 Minutes'" class="section-head-span">{{expiry.substring(0, 4)}}</span>
                <span v-else class="section-head-span">{{expiry.substring(0, 3)}}</span>
              </h5>
              <div class="table-responsive">
                <table class="table">
                  <thead>
                    <tr>
                      <td>Assets</td>
                      <td>Amount</td>
                      <td>Start Price</td>
                      <td>Close</td>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="(openBet, index) of openBets" :key="`openBet-${index}`">
                      <td v-if="openBet.position === 'buy'"><i class="fas fa-caret-up color-green"></i> BTC/DEMO</td>
                      <td v-if="openBet.position === 'sell'"><i style="transform: rotate(180deg);" class="fas fa-caret-up color-red"></i> BTC/DEMO</td>
                      <td>{{openBet.betsize}}</td>
                      <td>{{openBet.price}}</td>
                      <td class="text-center">{{openBet.minutes}}:{{openBet.seconds}}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
          <div class="col-lg-3 col-md-3">
            <div class="order-buy">
              <h5 class="section-head">
                The Book -
                <span v-if="expiry === '15 Minutes'" class="section-head-span">{{expiry.substring(0, 4)}}</span>
                <span v-else class="section-head-span">{{expiry.substring(0, 3)}}</span>
              </h5>
              <div class="table-responsive">
                <table class="table">
                  <thead>
                    <tr>
                      <td>Buy</td>
                      <td>Sell</td>
                    </tr>
                  </thead>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <div v-if="tradeViewLoading || !haveDemoBalance" class="trade-view-cover">
      <h4>Fetching data, almost there!</h4>
    </div>

  </div>
</template>

<script>
  import { HTTP } from '../http-constants'
  import router from '../router'
  import Chart from './Chart'
  var $ = require('jquery')
  window.jQuery = $
  var moment = require('moment')

  export default {
    name: 'trade',
    components: {
      Chart
    },
    props: ['auth', 'authenticated'],
    data () {
      return {
        tradeViewLoading: true,
        haveDemoBalance: false,
        demoBalance: null,
        betsize: 10.0,
        currentprice: null,
        expiry: '1 Minute',
        position: '',
        minutes: 0,
        seconds: 0,
        result: '',
        addtoExpiry: 0,
        timout: 0,
        multiplier: 0,
        auth0UUID: '',
        percentage: '',
        openBets: [],
        doneBets: [],
        // for/from the server
        simpleData: [],
        stringBets: '',
        stringBetsAsJson: [
          // {'pair': 'btc/demo', 'position': 'buy', 'price': 2.5, 'expiryint': 1, 'endsat': '6666666666666', 'betsize': 10, 'result': '', 'returns': 0, 'minutes': 0, 'seconds': 30}
        ]
      }
    },
    computed: {
      potentialProfit: function () {
        let potential
        if (this.expiry === '1 Minute') {
          potential = this.betsize * 0.8
          this.percentage = '+80%'
        } else if (this.expiry === '5 Minutes') {
          potential = this.betsize * 0.85
          this.percentage = '+85%'
        } else if (this.expiry === '15 Minutes') {
          potential = this.betsize * 0.9
          this.percentage = '+90%'
        }
        return potential
      }
    },
    methods: {
      getMoreDemo: function () {
        this.demoBalance = 100
        this.updateUserBalance()
      },
      chartLoaded: function (bool) {
        this.tradeViewLoading = false
      },
      getSimpleData: function () {
        let self = this
        setInterval(function () {
          jQuery.noConflict()
          $.getJSON('https://min-api.cryptocompare.com/data/histominute?fsym=BTC&tsym=USD&limit=9999&aggregate=1&e=CCCAGG&allData=true', function (data) {
            let simpleData = []
            data.Data.forEach(function (d) {
              let simpleDataItem = []
              simpleDataItem[0] = d.time
              simpleDataItem[1] = d.close
              simpleData.push(simpleDataItem)
            })
            self.simpleData = simpleData
          })
        }, 10000)
      },
      getCurrentPrice: function () {
        let self = this
        setInterval(function () {
          jQuery.noConflict()
          $.getJSON('https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD', function (data) {
            self.currentprice = data.USD
          })
        }, 2000)
      },
      getInitialPrice: function () {
        let self = this
        jQuery.noConflict()
        $.getJSON('https://min-api.cryptocompare.com/data/histominute?fsym=BTC&tsym=USD&limit=9999&aggregate=1&e=CCCAGG&allData=true', function (data) {
          let simpleData = []
          data.Data.forEach(function (d) {
            let simpleDataItem = []
            simpleDataItem[0] = d.time
            simpleDataItem[1] = d.close
            simpleData.push(simpleDataItem)
          })
          self.simpleData = simpleData
          self.getUserInfo()
          console.log(self.demoBalance)
        })
      },
      newBet: function (position, betsize = this.betsize, endingPrice = this.endingPrice, expiry = this.expiry, auth0UUID = this.auth0UUID) {
        if (this.demoBalance > 0 && this.demoBalance !== null && this.demoBalance >= betsize) {
          let self = this
          // GET STARTING PRICE FOR COMPARISON
          let startPrice = parseFloat(self.currentprice)
          self.position = position
          let expiryInt = 0

          self.demoBalance -= betsize
          // demoBalance -= betsize

          self.$emit('updateBalanceFromTrade', self.demoBalance)

          // check expiry
          if (expiry === '1 Minute') {
            self.timout = 60000
            self.multiplier = 0.8
            self.addtoExpiry = 1
            expiryInt = 1
          } else if (expiry === '5 Minutes') {
            self.timout = 60000 * 5
            self.multiplier = 0.85
            self.addtoExpiry = 5
            expiryInt = 5
          } else if (expiry === '15 Minutes') {
            self.timout = 60000 * 15
            self.multiplier = 0.9
            self.addtoExpiry = 15
            expiryInt = 15
          }
          // BET CONSTRUCTOR * BET * CONSTRUCTOR BET * CONSTRUCTOR
          let bet = {
            position: position,
            pair: 'BTC/USD',
            betsize: betsize,
            begunat: 0,
            endsat: 0,
            price: startPrice,
            endPrice: startPrice,
            expiryint: expiryInt,
            minutes: 0,
            seconds: 0,
            result: 'pending',
            returns: 0,
            updateBalance: this.updateUserBalance,
            // DISPLAY COUNTDOWN
            timout: self.timout,
            multiplier: self.multiplier,
            duration: 0,
            countDown: setInterval(function () {
              bet.duration = moment.duration(bet.duration.asMilliseconds() - 1000, 'milliseconds')
              bet.minutes = moment.duration(bet.duration).minutes()
              bet.seconds = moment.duration(bet.duration).seconds()
              if (bet.minutes === 0 && bet.seconds === 0) {
                clearInterval(bet.countDown)
              }
            }, 1000),
            // AS LONG AS COUNTDOWN
            /* eslint-disable */
            judge: setTimeout(function () {

              let removeFromOpenBets = function () {
                // console.log('removing at index: ' + self.openBets.indexOf(bet))
                let openBetIndex = self.openBets.indexOf(bet)
                if (openBetIndex > -1) {
                  self.openBets.splice(openBetIndex, 1)// REMOVE FROM PENDING BETS
                }
              }
              if (bet.position === 'buy') { // BUY
                bet.endPrice = parseFloat(self.currentprice) // GET CURRENT PRICE TO CHECK AGAINST
                if (bet.endPrice > bet.price) { // IF WIN
                  bet.returns = '+' + bet.betsize * bet.multiplier
                  // demoBalance += parseFloat(bet.betsize)
                  self.demoBalance += parseFloat(bet.betsize)
                  // demoBalance += parseFloat(bet.returns)
                  self.demoBalance += parseFloat(bet.returns)
                  bet.result = 'Won'
                  bet.updateBalance() // SEND BALANCE TO DB AND APP.vue
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
                else if (bet.endPrice === bet.price) { // IF TIED
                  // demoBalance += parseFloat(bet.betsize)
                  self.demoBalance += parseFloat(bet.betsize)
                  bet.result = 'Tied'
                  self.$emit('updateBalanceFromTrade', self.demoBalance)
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
                else { // IF LOST
                  bet.returns = bet.betsize * -1
                  bet.result = 'Lost'
                  bet.updateBalance()
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
              }
              else if (position === 'sell') { // SELL
                bet.endPrice = parseFloat(self.currentprice) // GET CURRENT PRICE TO CHECK AGAINST

                if (bet.endPrice < bet.price) { // IF WIN
                  bet.returns = '+' + bet.betsize * bet.multiplier
                  // demoBalance += parseFloat(bet.betsize)
                  self.demoBalance += parseFloat(bet.betsize)
                  // demoBalance += parseFloat(bet.returns)
                  self.demoBalance += parseFloat(bet.returns)
                  bet.result = 'Won'
                  bet.updateBalance()
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
                else if (bet.endPrice === bet.price) { // IF TIED
                  // demoBalance += parseFloat(bet.betsize)
                  self.demoBalance += parseFloat(bet.betsize)
                  bet.result = 'Tied'
                  self.$emit('updateBalanceFromTrade', self.demoBalance)
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
                else { // IF LOST
                  bet.returns = bet.betsize * -1
                  bet.result = 'Lost'
                  bet.updateBalance()
                  self.doneBets.unshift(bet) // PUSH TO DONE BETS
                  removeFromOpenBets() // REMOVE FROM OPEN BETS
                }
              }
            }, self.timout + 1000)
            /* eslint-enable */
          }
          // init
          bet.timout = self.timout
          bet.multiplier = self.multiplier
          bet.addtoExpiry = self.addtoExpiry

          // DETERMINE TIMESTAMPS
          let currentTime = moment()
          let expiryEnd = moment(currentTime).add(self.addtoExpiry, 'm')
          let diffTime = expiryEnd - currentTime
          let duration = moment.duration(diffTime, 'milliseconds')
          bet.duration = duration
          bet.begunat = currentTime
          bet.endsat = expiryEnd

          // push bet into list
          this.openBets.push(bet)
          this.stringBetsAsJson.push(bet)
          this.stringBets = JSON.stringify(this.stringBetsAsJson).replace(/"/g, "'").replace('[', '{').replace(/.$/, '}')
          this.updateUserBalance()
        }
      },
      // Get DemoBalance & Bets from db
      getUserInfo: function () {
        let self = this
        HTTP.get(`/api/userinfo/${self.auth0UUID}`)
          .then(response => {
            // Update balance
            self.demoBalance = response.data.DemoBalance
            console.log(self.demoBalance)
            self.haveDemoBalance = true

            self.$emit('updateBalanceFromTrade', response.data.DemoBalance)

            // DB has bets?
            if (response.data.StringBets) {
              self.stringBets = response.data.StringBets
              self.stringBetsAsJson = JSON.parse(self.stringBets.replace(/'/g, '"').replace('{', '[').replace(/.$/, ']'))
            } else {
              self.stringBets = JSON.stringify(self.stringBetsAsJson).replace(/"/g, "'").replace('[', '{').replace(/.$/, '}')
              console.log('first time? (no bets in db)', self.stringBets, self.stringBetsAsJson)
            }

            self.stringBetsAsJson.forEach((betFromDB, index) => { // LOOP THROUGH BETS FROM DB TO DETERMINE WHAT TABLE TO PLACE INTO
              if (betFromDB.pair) { // If valid bet
                let betEndsAtUnix = moment(betFromDB.endsat).unix()
                let endingPrice = null

                let rightNow = moment().format()
                let betEndsAt = moment(betFromDB.endsat).format()
                betFromDB.updateBalance = self.updateUserBalance

                if (betEndsAt > rightNow) { // CHECK IF BET IS STILL OPEN
                  if (betFromDB.expiryint === 1) { // check expiry
                    self.timout = 60000
                    self.multiplier = 0.8
                    self.addtoExpiry = 1
                  } else if (betFromDB.expiryint === 5) {
                    self.timout = 60000 * 5
                    self.multiplier = 0.85
                    self.addtoExpiry = 5
                  } else if (betFromDB.expiryint === 15) {
                    self.timout = 60000 * 15
                    self.multiplier = 0.9
                    self.addtoExpiry = 15
                  }

                  let currentTime = moment()
                  let endsAt = moment(betFromDB.endsat)
                  let newDiffTime = endsAt - currentTime
                  let newDuration = moment.duration(newDiffTime).asMilliseconds()
                  betFromDB.duration = newDuration

                  betFromDB.countDown = setInterval(function () { // COUNTDOWN
                    betFromDB.duration = moment.duration(betFromDB.duration, 'milliseconds') - 1000
                    betFromDB.minutes = moment.duration(betFromDB.duration).minutes()
                    betFromDB.seconds = moment.duration(betFromDB.duration).seconds()
                    // console.log(betFromDB.price, betFromDB.minutes, ':', betFromDB.seconds, 'dur:' + betFromDB.duration)
                    if (betFromDB.minutes === 0 && betFromDB.seconds === 0) {
                      clearInterval(betFromDB.countDown)
                    }
                  }, 1000)

                  betFromDB.judge = setTimeout(function () { // // // // // // // // // // // // // // THE JUDGE
                    let removeFromOpenBets = function () {
                      let openBetIndex = self.openBets.indexOf(betFromDB)
                      if (openBetIndex > -1) {
                        self.openBets.splice(openBetIndex, 1)// REMOVE FROM PENDING BETS
                      }
                    }
                    if (betFromDB.position === 'buy') { // BUY
                      betFromDB.endPrice = parseFloat(self.currentprice) // GET Ending PRICE TO CHECK AGAINST
                      console.log(betFromDB.endPrice)
                      if (betFromDB.endPrice > betFromDB.price) { // IF WIN
                        betFromDB.returns = '+' + betFromDB.betsize * betFromDB.multiplier
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        self.demoBalance += parseFloat(betFromDB.returns)
                        betFromDB.result = 'Won'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      } else if (betFromDB.endPrice === betFromDB.price) { // IF TIED
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        betFromDB.result = 'Tied'
                        betFromDB.updateBalance()
                        self.$emit('updateBalanceFromTrade', self.demoBalance)
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      } else { // IF LOST
                        betFromDB.returns = betFromDB.betsize * -1
                        betFromDB.result = 'Lost'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      }
                    } else if (betFromDB.position === 'sell') { // SELL
                      betFromDB.endPrice = parseFloat(self.currentprice) // GET Ending PRICE TO CHECK AGAINST
                      console.log(betFromDB.endPrice)
                      if (betFromDB.endPrice < betFromDB.price) { // IF WIN
                        betFromDB.returns = '+' + betFromDB.betsize * betFromDB.multiplier
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        self.demoBalance += parseFloat(betFromDB.returns)
                        betFromDB.result = 'Won'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      } else if (betFromDB.endPrice === betFromDB.price) { // IF TIED
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        betFromDB.result = 'Tied'
                        betFromDB.updateBalance()
                        self.$emit('updateBalanceFromTrade', self.demoBalance)
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      } else { // IF LOST
                        betFromDB.returns = betFromDB.betsize * -1
                        betFromDB.result = 'Lost'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                        removeFromOpenBets() // REMOVE FROM OPEN BETS
                      }
                    }
                  }, betFromDB.duration + 1000)

                  betFromDB.timout = self.timout
                  betFromDB.multiplier = self.multiplier
                  betFromDB.addtoExpiry = self.addtoExpiry

                  self.openBets.push(betFromDB)
                } else { // OR EXPIRED
                  if (betFromDB.result === 'pending') {
                    self.simpleData.forEach(d => {
                      let roundUp = betEndsAtUnix + 30
                      let roundDown = betEndsAtUnix - 30
                      if (d[0] > roundDown && d[0] < roundUp) {
                        endingPrice = d[1]
                        console.log(self.simpleData)
                        console.log(endingPrice)
                        console.log(roundUp, roundDown, betEndsAtUnix)
                      }
                    })
                    if (betFromDB.position === 'buy') { // BUY
                      betFromDB.endPrice = parseFloat(endingPrice) // GET Ending PRICE TO CHECK AGAINST
                      console.log(betFromDB.endPrice)
                      if (betFromDB.endPrice > betFromDB.price) { // IF WIN
                        betFromDB.returns = '+' + betFromDB.betsize * betFromDB.multiplier
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        self.demoBalance += parseFloat(betFromDB.returns)
                        betFromDB.result = 'Won'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      } else if (betFromDB.endPrice === betFromDB.price) { // IF TIED
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        betFromDB.result = 'Tied'
                        betFromDB.updateBalance()
                        self.$emit('updateBalanceFromTrade', self.demoBalance)
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      } else { // IF LOST
                        betFromDB.returns = betFromDB.betsize * -1
                        betFromDB.result = 'Lost'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      }
                    } else if (betFromDB.position === 'sell') { // SELL
                      betFromDB.endPrice = parseFloat(endingPrice) // GET Ending PRICE TO CHECK AGAINST
                      console.log(betFromDB.endPrice)
                      if (betFromDB.endPrice < betFromDB.price) { // IF WIN
                        betFromDB.returns = '+' + betFromDB.betsize * betFromDB.multiplier
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        self.demoBalance += parseFloat(betFromDB.returns)
                        betFromDB.result = 'Won'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      } else if (betFromDB.endPrice === betFromDB.price) { // IF TIED
                        self.demoBalance += parseFloat(betFromDB.betsize)
                        betFromDB.result = 'Tied'
                        betFromDB.updateBalance()
                        self.$emit('updateBalanceFromTrade', self.demoBalance)
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      } else { // IF LOST
                        betFromDB.returns = betFromDB.betsize * -1
                        betFromDB.result = 'Lost'
                        betFromDB.updateBalance()
                        self.doneBets.unshift(betFromDB) // PUSH TO DONE BETS
                      }
                    }
                  } else {
                    self.doneBets.unshift(betFromDB)
                  }
                }
              } else {
                console.log('betFromDB pair: ', 'no valid bets provided')
              }
            })
          })
          .catch(e => {
            console.log('err: ' + e)
          })
      },
      updateUserBalance: function () { // Happens after each bet
        this.$emit('updateBalanceFromTrade', this.demoBalance)
        this.stringBets = JSON.stringify(this.stringBetsAsJson).replace(/"/g, "'").replace('[', '{').replace(/.$/, '}')

        HTTP.post(`/api/userinfo/${this.auth0UUID}`, {
          'demo_balance': this.demoBalance,
          'stringbets': this.stringBets
        }).then(function (response) {})
          .catch(function (error) {
            this.errors = error
            console.log(error)
          })
      }
    },
    mounted () {
      this.auth0UUID = localStorage.getItem('auth0UUID')
      // console.log('TRADE(mount): auth0uuid: ', this.auth0UUID)
      if (this.auth0UUID) {
        this.getInitialPrice()
        this.getSimpleData()
        this.getCurrentPrice()
      }
    },
    beforeMount () {
      if (!this.authenticated) {
        router.replace('/home')
      }
    }
  }
</script>

<style lang="scss" scoped>
/*
  tr {
  	display: none;
  	td {
  		font-size: 12px;
  	}
  }
  tr:nth-child(-n+4) {
  	display: table-row;
  }
*/

tr {
	td {
		font-size: 12px;
	}
}

.trade-view-cover {
	background-color: #051E2C;
	bottom: 0;
	display: block;
	left: 0;
	position: absolute;
	right: 0;
	top: 0;
	z-index: 9999;

	h4 {
		color: white;
		margin-top: 30%;
		text-align: center;
	}
}

#main-wrapper {
	background-image: '../assets/images/bj.jpg';
}

.trade-page {
	&__header {
		align-items: center;
		background-color: rgb(135, 162, 235);
		display: flex;
		justify-content: space-between;

		.logo,
		.demo-balance {
			margin: 0;
		}
	}
}

.nav-link:hover {
	i {
		color: #01D014 !important;
	}
}

.deposit {
	a,
	i {
		color: white !important;
	}
}
.deposit:hover {
	a,
	i {
		background-color: white;
		color: #1E84E0 !important;
	}
}

.section-head {
	text-transform: initial !important;
}
.trades-section,
.order-pending,
.order-buy {
	min-height: 296px;
}

.table-responsive {
	max-height: 216px;
	overflow-y: scroll;
}
</style>

