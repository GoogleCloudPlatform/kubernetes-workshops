# Copyright 2016 Google, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ruby:2.3

WORKDIR /app

COPY Gemfile Gemfile.lock ./
RUN echo 'gem "sqlite3"' >> Gemfile
RUN echo 'gem "therubyracer"' >> Gemfile
RUN bundle install --without development test

COPY . ./
RUN echo 'gem "sqlite3"' >> Gemfile
RUN echo 'gem "therubyracer"' >> Gemfile

COPY database.yml ./config/
RUN rake db:schema:load

RUN echo "Lobsters::Application.config.secret_key_base = '$(rake secret)'" > config/initializers/secret_token.rb

COPY production.rb ./config/initializers/
RUN rake db:seed

EXPOSE 3000
ENV RAILS_ENV=development
CMD ["/usr/local/bundle/bin/rails", "server"]
