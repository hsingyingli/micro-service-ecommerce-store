import React, { Fragment } from "react";
import { Menu, Transition } from '@headlessui/react'
import { useAuth } from "@/hooks/useAuth";
import { UserIcon } from '@heroicons/react/24/outline'
import Link from "next/link";
import { Tooltip } from "../Tooltip";

const AccountMenu: React.FC = () => {
  const { user, logout } = useAuth()
  return (
    <div>
      <Menu as="div" className="relative inline-block text-left">
        <div>
          <Tooltip tip="Account">
            <Menu.Button className="inline-flex w-full justify-center 
            rounded-md bg-secondary-400 p-2 hover:bg-secondary-500 transition-colors duration-150">
              <UserIcon
                className="h-5 w-5 text-violet-200 hover:text-violet-100"
                aria-hidden="true"
              />
            </Menu.Button>
          </Tooltip>
        </div>
        <Transition
          as={Fragment}
          enter="transition ease-out duration-100"
          enterFrom="transform opacity-0 scale-95"
          enterTo="transform opacity-100 scale-100"
          leave="transition ease-in duration-75"
          leaveFrom="transform opacity-100 scale-100"
          leaveTo="transform opacity-0 scale-95"
        >
          <Menu.Items className="absolute right-0 mt-2 w-56 origin-top-right divide-y divide-gray-100 
            rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
            <div className="p-1 ">
              {
                user === null ?
                  <Menu.Item>
                    <Link
                      href={"/login"}
                      className="flex w-full items-center rounded-md p-2 
                      hover:bg-secondary-50"
                    >
                      Login
                    </Link>
                  </Menu.Item>
                  : (
                    <>
                      <Menu.Item>
                        <Link
                          href={"/Profile"}
                          className="flex w-full items-center rounded-md p-2 
                      hover:bg-secondary-50"
                        >
                          Profile
                        </Link>
                      </Menu.Item>
                      <Menu.Item>
                        <button
                          onClick={logout}
                          className="flex w-full items-center rounded-md p-2 
                      hover:bg-secondary-50"
                        >
                          Logout
                        </button>
                      </Menu.Item>
                    </>
                  )
              }
            </div>
          </Menu.Items>
        </Transition>
      </Menu>
    </div>
  )
}

export {
  AccountMenu
}
